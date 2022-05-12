import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.scss';
import App from './App';
import reportWebVitals from './reportWebVitals';
import init from './init';

const app = init();
const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(app);


// const container = document.getElementById('root');
// ReactDOM.render(app, container);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
