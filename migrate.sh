#!/bin/bash
docker run --rm -v $(pwd):/www -v /data:/data gobuffalo/buffalo:v0.10.1 sh -c "cd /www;buffalo db migrate up"