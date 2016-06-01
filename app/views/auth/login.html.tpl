<div class="container">
    <div class="row">
        <div class="column column-12">
            <div class="auth login">
                <form action="{{mountpathed "login"}}" method="POST">
                    <div>
                      Demo Account: dev@getqor.com/testing
                    </div>

                    <ul class="auth-form">
                        {{if .error}}
                            <li class="error">{{.error}}</li>
                        {{end}}
                        <li>
                            <label for="{{.primaryID}}">{{ t "login.email" }}</label>
                            <input type="text" class="form-control" name="{{.primaryID}}" placeholder="{{title .primaryID}}" value="{{.primaryIDValue}}">
                        </li>

                        <li {{if .showRecover}}class="forget-password"{{end}}>
                            <label for="{{.primaryID}}">{{ t "login.password" }}</label>
                            {{if .showRecover}}<a href="{{mountpathed "recover"}}" class="forget-password-link">{{ t "login.forget_password" }}</a>{{end}}
                            <input type="password" class="form-control" name="password" placeholder="Password"><br />
                            <input type="hidden" name="{{.xsrfName}}" value="{{.xsrfToken}}" />
                        </li>


                        {{if .showRemember}}<li><label><input type="checkbox" name="rm" value="true"> {{ t "login.remember_me" }}</label></li>{{end}}


                        <li>
                            <button type="submit" class="button button__primary">{{ t "login.button" }}</button>
                        </li>

                        {{if .showRegister}}<li>{{ t "login.Need an Account?" }} <a href="{{mountpathed "register"}}">{{ t "login.register" }}</a></li>{{end}}
                    </ul>

                </form>
            </div>
        </div>
    </div>
</div>
