import * as React from 'react';
import styled from 'styled-components/macro';
import { Helmet } from 'react-helmet-async';

export function NotFoundPage() {
  return (
    <>
      <Helmet>
        <title>404 Page Not Found</title>
        <meta name="description" content="Page not found" />
      </Helmet>
      <Wrapper>
        <Title>
          4
          <span role="img" aria-label="Crying Face">
            😢
          </span>
          4
        </Title>
        <b>Page not found.</b>
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
