<style>
  form {
    width: 200px;
    margin: auto;
    margin-top: 20px;
  }
  button {
    margin-top: 10px;
    border: none;
    border-radius: 5px;
    width: 80px;
    background-color: rgb(48, 166, 221);
    color: white;
    margin-bottom: 10px;
  }
  input {
    margin-bottom: 10px;
  }
</style>

<form action="{{mountpathed "recover"}}" method="POST">
  <label style="margin-bottom:0px;" for="{{.primaryID}}">Email:</label>
  <input type="text" name="{{.primaryID}}" placeholder="{{title .primaryID}}" value="{{.primaryIDValue}}" /><br />
  {{$pid := .primaryID}}{{with .errs}}{{with $errlist := index . $pid}}{{range $errlist}}<span>{{.}}</span><br />{{end}}{{end}}{{end}}
  <label style="margin-bottom:0px;" for="{{.primaryID}}">Email Confirmation:</label>
  <input type="text" name="confirm_{{.primaryID}}" placeholder="Confirm {{title .primaryID}}" value="{{.confirmPrimaryIDValue}}" /><br />
  {{$cpid := .primaryID | printf "confirm_%s"}}{{with .errs}}{{with $errlist := index . $cpid}}{{range $errlist}}<span>{{.}}</span><br />{{end}}{{end}}{{end}}
  <input type="hidden" name="{{.xsrfName}}" value="{{.xsrfToken}}" />
  <button type="submit">Recover</button><br />
</form>
