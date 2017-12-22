import Header from './Header';
import Footer from './Footer';

const MainLayout = props => (
    <div>
        <Header />
        {props.children}
        <Footer />
    </div>
);

export default MainLayout;
