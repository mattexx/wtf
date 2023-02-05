# wtf

it's kind of like a file browser, except it doesn't browse

## dev

1. Put more images in `img/`
   - Rules
     - Only one `.` in the image filenames (no multi-extensions)
     - Extension must map to a valid `image/*` content type
   - Guidelines
     - kebab-case-file-names
2. `docker build . -t wtf`
3. `docker run -v "$(pwd)/img:/app/img/:ro" -p 8080:8080 wtf`
4. [http://localhost:8080/darkside](http://localhost:8080/darkside)

## run

`docker run -d --restart=unless-stopped --name wtf -v "$(pwd)/img:/app/img/:ro" -p 8080:8080 wtf`

## demo

[http://texx.wtf](http://texx.wtf)
