# wtf

it's kind of like a file browser, except it doesn't browse

## dev

1. `mkdir img`
2. Put images in `img/`
   - Rules
     - Only one `.` in the image filenames (no multi-extensions)
     - Extension must map to a valid `image/*` content type
   - Guidelines
     - kebab-case-file-names
3. `docker build . -t wtf`
4. `docker run -v "$(pwd)/img:/app/img/:ro" -p 8080:8080 wtf`
