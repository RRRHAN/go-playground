import * as React from 'react';
import styled from 'styled-components/macro';
import { Helmet } from 'react-helmet-async';

export function HomePage() {
  return (
    <>
      <Helmet>
        <title>Home Page</title>
        <meta name="description" content="A Boilerplate application homepage" />
      </Helmet>
      <Wrapper>
        <Title>
          4
          <span role="img" aria-label="Smiling Face">
            😁
          </span>
          4
        </Title>
        <b>Dashboard Page</b>
      </Wrapper>
    </>
  );
}

const Wrapper = styled.div.attrs({
  className:
    'w-full h-screen bg-gray-100 p-2 flex items-center justify-center flex-col',
})``;

const Title = styled.div.attrs({
  className: 'text-black text-xl text-center block mx-auto',
})``;
