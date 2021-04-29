# intervals

Backend for BSRs super simple planning and tracking tool.

## Development

Sample VSCode launch config

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}",
      "env": {
        "PORT": "8080",
        "DATABASE_URL": "postgres://postgres:[password]@[host]:[port]/[db]"
      },
      "args": []
    }
  ]
}
```

You can tail logs from codespaces using heroku CLI.

```bash
$ curl https://cli-assets.heroku.com/install.sh | sh
$ heroku login -i
$ heroku logs --tail --app intervals-bsr
```

### Sample requests

Create an athlete
```bash
curl -d '{"name":"Sample", "category":"D"}' -H "Content-Type: application/json" -X POST http://localhost:8080/athletes
```

Retrieve athletes
```bash
curl  -H "Content-Type: application/json" -X GET http://localhost:8080/athletes
```