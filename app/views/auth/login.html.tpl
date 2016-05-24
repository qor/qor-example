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

<form action="{{mountpathed "login"}}" method="POST">
    {{if .error}}{{.error}}<br />{{end}}
	<label style="margin-bottom:0px;" for="{{.primaryID}}">{{ t "login.email" }}</label>
    <input type="text" class="form-control" name="{{.primaryID}}" placeholder="{{title .primaryID}}" value="{{.primaryIDValue}}"><br />
	<label style="margin-bottom:0px;" for="{{.primaryID}}">{{ t "login.password" }}</label>
    <input type="password" class="form-control" name="password" placeholder="Password"><br />
    <input type="hidden" name="{{.xsrfName}}" value="{{.xsrfToken}}" />
    {{if .showRemember}}<input type="checkbox" name="rm" value="true"> {{ t "login.remember_me" }}{{end}}
    <button type="submit">{{ t "login.button" }}</button><br />
    {{if .showRecover}}<a href="{{mountpathed "recover"}}">{{ t "login.forget_password" }}</a>{{end}}
    {{if .showRegister}}<a href="{{mountpathed "register"}}">{{ t "login.register" }}</a>{{end}}
</form>
