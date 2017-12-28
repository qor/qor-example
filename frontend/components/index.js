import Link from 'next/link';
require('es6-promise').polyfill();
require('isomorphic-fetch');

const widgetPrefix = 'http://localhost:7000/admin/page_builder_widgets/';
const widgetNames = ['home page banner', 'men collection', 'women collection', 'new arrivals promotion', 'model products'];

async function asyncForEach(array, callback) {
    for (let index = 0; index < array.length; index++) {
        await callback(array[index], index, array);
    }
}

let initialProps = async () => {
    let widgetData = {};

    await asyncForEach(widgetNames, async widget => {
        let name = widget.replace(/\s/g, '');

        widget = await fetch(`${widgetPrefix}${widget}.json`);
        widget = await widget.json();
        widgetData[name] = widget;
    });

    return widgetData;
};

export default initialProps;
