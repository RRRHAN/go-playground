import * as React from 'react';
import { Helmet } from 'react-helmet-async';

const foo = String(process.env.REACT_APP_FOO);

export function Foo() {
  return (
    <>
      <Helmet>
        <title>Foo</title>
      </Helmet>
      <h1 className="text-center text-3xl">FOO : {foo}</h1>
    </>
  );
}
