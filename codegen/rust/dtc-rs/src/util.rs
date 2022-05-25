use crate::error::Error;

pub struct Framed {
    buf: Vec<u8>,
}

impl Framed {
    pub fn new() -> Self {
        Self { buf: Vec::new() }
    }

    pub fn push<T: Fn(&[u8])>(&mut self, src: &[u8], f: T) {
        let mut data = src.as_ref();
        if data.len() == 0 {
            return;
        }
        // Try
        if self.buf.is_empty() {
            match self.process(src, f) {
                Ok(Some(idx)) => {
                    data = &data[idx..];
                    if data.len() > 0 {
                        self.buf.extend_from_slice(data);
                    }
                }
                Ok(None) => {}
                Err(_e) => {}
            }
        } else {
            let size = if self.buf.len() > 1 {
                (unsafe { u16::from_le(*(self.buf.as_ptr() as *const u16)) }) as usize
            } else {
                let b: [u8; 2] = [self.buf.as_slice()[0], src.as_ref()[0]];
                (unsafe { u16::from_le(*(b.as_ptr() as *const u16)) }) as usize
            };

            // Is next message received?
            if (self.buf.len() + src.len()) < size {
                self.buf.extend_from_slice(src.as_ref());
                return;
            }

            let split_at = size - self.buf.len();
            self.buf.extend_from_slice(&data[0..split_at]);
            data = &data[split_at..];
            f(self.buf.as_slice());
            self.buf.clear();

            match self.process(data, f) {
                Ok(Some(idx)) => {
                    data = &data[idx..];
                    if data.len() > 0 {
                        self.buf.extend_from_slice(data);
                    }
                }
                Ok(None) => {}
                Err(_e) => {}
            }

            if data.len() < 4 {
                self.buf.extend_from_slice(data);
            }
        }
    }

    fn process<T: Fn(&[u8])>(&mut self, mut data: &[u8], f: T) -> Result<Option<usize>, Error> {
        let mut size = unsafe { u16::from_le(*(data.as_ptr() as *const u16)) } as usize;
        // Bad message?
        if size < 4 {
            return Err(Error::Malformed("bad size".into()));
        }
        if size > data.len() {
            // buf.extend_from_slice(data);
            return Ok(Some(0usize));
        }

        if size == data.len() {
            f(&data[..size]);
            Ok(None)
        } else {
            let mut offset = 0;
            loop {
                f(&data[..size]);

                offset += size;
                data = &data[size..];

                if data.len() == 0 {
                    return Ok(None);
                } else if data.len() < 2 {
                    return Ok(Some(offset));
                }
                size = unsafe { u16::from_le(*(data.as_ptr() as *const u16)) } as usize;
                if size > data.len() {
                    return Ok(Some(offset));
                }
            }
        }
    }
}
