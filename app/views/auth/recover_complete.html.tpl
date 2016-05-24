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

<form action="{{mountpathed "recover/complete"}}" method="POST">
  <label style="margin-bottom:0px;" for="{{.primaryID}}">{{ t "recover_complete.new_password" }}</label>
  <input type="hidden" name="token" value="{{.token}}" />
  <input type="password" name="password" value="" /><br />
  {{with .errs}}{{with $errlist := index . "password"}}{{range $errlist}}<span>{{.}}</span><br />{{end}}{{end}}{{end}}
  <label style="margin-bottom:0px;" for="{{.primaryID}}">{{ t "recover_complete.confirm_new_password" }}</label>
  <input type="password" name="confirm_password" value="" /><br />
  {{with .errs}}{{with $errlist := index . "confirm_password"}}{{range $errlist}}<span>{{.}}</span><br />{{end}}{{end}}{{end}}
  <input type="hidden" name="{{.xsrfName}}" value="{{.xsrfToken}}" />
  <button type="submit">{{ t "recover_complete.button" }}</button><br />
</form>
