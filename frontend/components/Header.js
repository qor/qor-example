import React, {Component} from 'react';
import Link from 'next/link';
require('es6-promise').polyfill();
require('isomorphic-fetch');

class Header extends Component {
    constructor() {
        super();
        this.state = {
            categories: []
        };
    }
    componentDidMount() {
        fetch('http://localhost:7000/admin/categories.json')
            .then(res => res.json())
            .then(res => {
                this.setState({
                    categories: res
                });
            });
    }
    render() {
        const menus = this.state.categories.map(category => {
            return (
                <li key={category.ID}>
                    <Link href={`/category/${category.Name}`}>
                        <a>{category.Name}</a>
                    </Link>
                </li>
            );
        });

        return (
            <div>
                <Link href="/">
                    <a>
                        <img src="/static/images/logo.png" width="112" />
                    </a>
                </Link>
                <ul>{menus}</ul>
                <ul>
                    <li>
                        <Link href="/account">
                            <a>MY ACCOUNT</a>
                        </Link>
                    </li>
                    <li>
                        <Link href="/cart">
                            <a>CART</a>
                        </Link>
                    </li>
                </ul>
            </div>
        );
    }
}

export default Header;
