package workspace

var CreateWksp = `{
    "name": "ts-poc-job",
    "type": [
        "NOT_SET"
    ],
    "resource_group": "Default",
    "description": "terraform workspace",
    "template_repo": {
        "url": "https://github.com/tanya-shanker/simple-terraform",
        "branch": ""
    },
    "template_data": [
        {
            "folder": ".",
            "type": "terraform_v0.13"
        }
    ]
}`
