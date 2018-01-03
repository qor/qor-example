import React, {Component} from 'react';
import Link from 'next/link';
import Head from 'next/head';

require('es6-promise').polyfill();
require('isomorphic-fetch');

const AUTHAPI = 'http://localhost:7000/auth/password/login';

class Login extends Component {
    constructor(props) {
        super(props);
        this.state = {
            login: '',
            password: ''
        };
    }

    setStateAsync(state) {
        return new Promise(resolve => {
            this.setState(state, resolve);
        });
    }

    handleChange(e) {
        let state = {};
        state[e.target.name] = e.target.value;
        console.log(state);
        this.setState(state);
    }
    async handleSubmitLogin(e) {
        e.preventDefault();
        const res = await fetch(AUTHAPI, {
            method: 'POST',
            headers: {
                Accept: 'application/json, text/plain, */*',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                login: this.state.login,
                password: this.state.password
            })
        });
        const response = await res.json();

        // await this.setStateAsync({ipAddress: ip})
        console.log(response);
    }
    render() {
        return (
            <div>
                <Head>
                    <link rel="stylesheet" href="/static/stylesheets/qor_auth.css" />
                </Head>
                <div>
                    <div className="container qor-auth">
                        <div className="qor-auth-box">
                            <h1>Sign In</h1>

                            <a className="signup-link" href="/auth/register">
                                Don't have an account? Click here to sign up.
                            </a>

                            <form onSubmit={e => this.handleSubmitLogin(e)}>
                                <strong>Demo Account: dev@getqor.com / testing</strong>
                                <ul className="auth-form">
                                    <li>
                                        <label>Email</label>
                                        <input type="email" id="email" name="login" placeholder="email address" value={this.state.email} onChange={e => this.handleChange(e)} />
                                    </li>
                                    <li>
                                        <label>Password</label>
                                        <a className="forgot-password" href="/auth/password/new">
                                            Forgot Password?
                                        </a>
                                        <input type="password" className="form-control" id="password" name="password" placeholder="Password" onChange={e => this.handleChange(e)} />
                                    </li>
                                    <li>
                                        <button type="submit" className="button button__primary">
                                            Sign in
                                        </button>
                                    </li>
                                </ul>
                            </form>

                            <div className="line">
                                <span>OR SIGN IN WITH</span>
                            </div>

                            <div className="qor-auth-social-login">
                                <a href="/auth/facebook/login" className="qor-auth-facebook" title="Sign in with facebook">
                                    <i className="fa fa-facebook" aria-hidden="true" />
                                </a>

                                <a href="/auth/twitter/login" className="qor-auth-twitter" title="Sign in with twitter">
                                    <i className="fa fa-twitter" aria-hidden="true" />
                                </a>

                                <a href="/auth/github/login" className="qor-auth-github" title="Sign in with github">
                                    <i className="fa fa-github" aria-hidden="true" />
                                </a>

                                <a href="/auth/google/login" className="qor-auth-google" title="Sign in with google">
                                    <i className="fa fa-google" aria-hidden="true" />
                                </a>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

export default Login;
