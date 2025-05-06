# This is a for fun project, don't expose it to the internet

With that out of the way, I felt like learning basic backend development in Go to pad my resume.
I'm not a web developer, feel free to send pull requests if you find it useful and know how to improve it 

## How to use:

Create a new folder in the .config directory called gowebdeck with a scripts.json in it.
### Sample scripts.json:
```json 
[
    {
        "id": 1,
        "path": "/path/to/your/script",
        "description": "This script does something.",
        "icon": "path/to/icon.png"
    },
    {
        "id": 2,
        "path": "path/to/your/script",
        "description": "this other script does something else.",
        "icon": "path/to/icon2.png"
    },
    {
        "id": 3,
        "path": "path/to/your/script",
        "path": "path/to/your/script3",
        "description": "this script does yet another thing.",
        "icon": "path/to/icon3.png"
    },
    {
        "id": 4,
        "path": "path/to/your/script",
        "path": "path/to/your/script4",
        "description": "this script does something different.",
        "icon": "path/to/icon4.png"
    },
    {
        "id": 5,
        "path": "path/to/your/script",
        "path": "path/to/your/script5",
        "description": "this script does something unique.",
        "icon": "path/to/icon5.png"
    }
]
```
# TODO:
- [ ] Add support for more scripting languages
- [ ] Add system monitoring scripts
- [ ] Finish the web interface
- [ ] Add a way to create new scripts from the web interface (maybe)

