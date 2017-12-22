import React, {Component} from 'react';
import Link from 'next/link';
import 'whatwg-fetch';

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
            </div>
        );
    }
}

// const Header = props => (
//     <div>
//         <Link href="/">
//             <a>
//                 <img src="/static/images/logo.png" width="112" />
//             </a>
//         </Link>
//         <ul>
//             {props.propTypes}
//             {props.categories.map(({category}) => (
//                 <li key={category.ID}>
//                     <Link href={`/category/${category.Name}`}>
//                         <a>{category.Name}</a>
//                     </Link>
//                 </li>
//             ))}
//         </ul>
//     </div>
// );

// Header.getInitialProps = async ({req}) => {
//     const res = await fetch('http://localhost:7000/admin/categories.json');
//     const data = await res.json();

//     console.log(`Categories data fetched. Count: ${data.length}`);

//     return {
//         categories: data
//     };
// };

export default Header;
