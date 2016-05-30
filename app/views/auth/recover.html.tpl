<div class="container">
    <div class="row">
        <div class="column column-12">
            <div class="auth login">
                <form action="{{mountpathed "recover"}}" method="POST">
                    <ul class="auth-form">
                        <li>
                            <label style="margin-bottom:0px;" for="{{.primaryID}}">{{ t "recover.email" }}</label>
                            <input type="text" name="{{.primaryID}}" value="{{.primaryIDValue}}" />
                            {{$pid := .primaryID}}{{with .errs}}{{with $errlist := index . $pid}}{{range $errlist}}<p class="error">{{.}}</p>{{end}}{{end}}{{end}}
                        </li>

                        <li>
                            <label style="margin-bottom:0px;" for="{{.primaryID}}">{{ t "recover.confirm_email" }}</label>
                            <input type="text" name="confirm_{{.primaryID}}" value="{{.confirmPrimaryIDValue}}" />
                            {{$cpid := .primaryID | printf "confirm_%s"}}{{with .errs}}{{with $errlist := index . $cpid}}{{range $errlist}}<p class="error">{{.}}</p>{{end}}{{end}}{{end}}
                            <input type="hidden" name="{{.xsrfName}}" value="{{.xsrfToken}}" />
                        </li>
                        <li><button type="submit" class="button button__primary">{{ t "recover.button" }}</button></li>
                    </ul>
                </form>
            </div>
        </div>
    </div>
</div>
