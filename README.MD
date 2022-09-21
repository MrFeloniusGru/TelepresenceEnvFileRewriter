### Telepresence Environment variables file tool
*The tool* aim to remove environment variables from the *telepresence file --env-file* (https://www.telepresence.io/docs/v2.0/reference/environment/) these defined in a application settings file.

For build the tool docker command can be used in the repository root directory 
```
docker build --pull --rm -f "Dockerfile" -t docker_image_name:latest "."
```

#### Usage: ####

--env-file *intercept.env*:

```
MF_COMMUNICATION_LINE_PORT_27004_TCP_PROTO=tcp
MF_COMMUNICATION_LINE_SERVICE_HOST=10.44.128.155
MF_COMMUNICATION_LINE_SERVICE_PORT=27004
AUTH__HEADER__USERIDHEADERNAME=X-User-Id
```

*appsettings.json*:

```
{
 ...
  "Auth": {
    "Header": {
      "UserIdHeaderName": "X-User-Id",
      "PermissionHeaderName": "X-Permission",
      "BranchIdHeaderName": "X-Branch-Id",
      "SubBranchIdHeaderName": "X-SubBranch-Id"
    }
  }
}
```

command:
```
docker run -it --rm \
-v ./intercept.env:/intercept.env \
-v ./appsettings.json:/appsettings.json \
docker_image_name:latest \
-an /appsettings.Development.json -en /cl-intercept.env -at "dotnet-json"
```

The result execution the command is *intercept.env* without AUTH__HEADER__USERIDHEADERNAME Line
