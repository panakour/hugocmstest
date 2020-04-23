import React from 'react';
import {Datagrid, List, TextField} from 'react-admin';


const postRowClick = (id: any, basePath: any, record: any) => basePath + "/" + record.path;
const PageList = (props: any) => (
    <List {...props}>
        <Datagrid rowClick={postRowClick}>
            <TextField source="id"/>
            <TextField source="path"/>
            <TextField source="name"/>
            <TextField source="params.head_image"/>
            <TextField source="content"/>
        </Datagrid>
    </List>

);

export default PageList;
