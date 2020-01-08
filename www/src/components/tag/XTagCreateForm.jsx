import { useMutation } from '@apollo/react-hooks';
import gql from 'graphql-tag';
import React, { useState } from 'react';
import { Form, Loader } from 'semantic-ui-react';
import { MULTI_TAG_QUERY } from '../../views';
import { XErrorMessage } from '../messages';

const CREATE_TAG = gql`
mutation CreateTag($name: String!) {
    createTag(input: {name: $name} ) {
        id
        name
    }
}`;

const XTagCreateForm = () => {
    const [createTag, { loading, error }] = useMutation(CREATE_TAG, {
        refetchQueries: [{ query: MULTI_TAG_QUERY }],
    });

    const [params, setParams] = useState({ name: '' });
    const handleChange = (e, { name, value }) => {
        setParams({ ...params, [name]: value });
    }
    const handleSubmit = () => {
        console.log("Creating tag with params: ", params);
        createTag({ variables: params }).then(({ data }) => {
            console.log("Create tag result: ", data);
        });
    }

    return (
        <Form onSubmit={handleSubmit}>
            <Form.Input
                label={{ content: 'Name' }}
                labelPosition='right'
                placeholder='Enter tag name'
                name='name'
                value={params.name}
                onChange={handleChange}
            />

            <Form.Button positive floated='right'>Create</Form.Button>

            {error ?
                <XErrorMessage title='Failed to create tag' msg={`${error}`} />
                : <Loader active={loading} />}
        </Form>
    );
}

export default XTagCreateForm;