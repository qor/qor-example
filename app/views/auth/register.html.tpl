<div class="container">
    <div class="row">
        <div class="column column-12">
            <div class="auth register">
				<form action="{{mountpathed "register"}}" method="post">
					<ul class="auth-form">
						<li>
							<label for="{{.primaryID}}">{{ t "register.email" }}</label>
							<input name="{{.primaryID}}" type="text" value="{{with .primaryIDValue}}{{.}}{{end}}" placeholder="{{title .primaryID}}" />
							{{$pid := .primaryID}}{{with .errs}}{{with $errlist := index . $pid}}{{range $errlist}}<p class="error">{{.}}</p>{{end}}{{end}}{{end}}
						</li>

						<li>
							<label for="password">{{ t "register.password" }}</label>
							<input name="password" type="password" placeholder="Password" />
							{{with .errs}}{{range .password}}<p class="error">{{.}}</p>{{end}}{{end}}
						</li>

						<li>
							<label for="confirm_password">{{ t "register.confirm_password" }}:</label>
							<input name="confirm_password" type="password" />
							{{with .errs}}{{range .confirm_password}}<p class="error">{{.}}</p>{{end}}{{end}}
						</li>

					    <li><button type="submit" class="button button__primary">{{ t "register.button" }}</button></li>

						<input type="hidden" name="{{.xsrfName}}" value="{{.xsrfToken}}" />
					</ul>
				</form>
			</div>
		</div>
	</div>
</div>
