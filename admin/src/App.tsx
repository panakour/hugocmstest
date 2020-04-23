import React from 'react';
import './App.css';
import dataprovider from "./dataprovider";
import {Admin, Resource} from 'react-admin';
import PageList from "./pages/PageList";

const App = () => <Admin dataProvider={dataprovider("http://localhost:8080/api/v1/content/sections")}>
    <Resource name="project" list={PageList}/>
    <Resource name="contact" list={PageList}/>
</Admin>;

export default App;
