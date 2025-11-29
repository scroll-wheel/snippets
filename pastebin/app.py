from flask import Flask, request
from pathlib import Path

directory = 'uploads'
app = Flask(__name__, static_url_path='', static_folder=directory)
Path(directory).mkdir(exist_ok=True)

@app.put('/<filename>')
def put_file(filename):
    if filename == '.' or filename == '..':
        return {'status': 403, 'message': 'Forbidden'}
    path = Path(directory).joinpath(filename)
    path.write_bytes(request.get_data())
    return {'status': 200, 'message': 'OK'}


@app.route('/', methods=['GET'])
def upload_file():
    return f"""<pre>
PUT files here:
    curl -T &ltfile&gt {request.url_root}[&ltname&gt]
    powershell.exe iwr -me put -i &ltfile&gt {request.url_root}&ltname&gt
</pre>
"""