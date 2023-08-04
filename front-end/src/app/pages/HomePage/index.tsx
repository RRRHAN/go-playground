import * as React from 'react';
import { Helmet } from 'react-helmet-async';
import axios from 'axios';

const BASE_URL = String(process.env.REACT_APP_BASE_URL);

export function HomePage() {
  const [images, setImages] = React.useState([{ id: 0, name: '' }]);
  const [uploadedFiles, setUploadedFiles] = React.useState<FileList | null>(
    null,
  );

  const handleCahngeInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    e.preventDefault();
    setUploadedFiles(e.target.files);
  };

  const handleUploadFile = (
    e: React.MouseEvent<HTMLInputElement, MouseEvent>,
  ) => {
    if (uploadedFiles === null) return;
    const files = uploadedFiles as FileList;
    const formData = new FormData();
    for (let index = 0; index < files.length; index++) {
      formData.append('images', files[index]);
    }
    axios
      .post(BASE_URL + 'images', formData)
      .then(response => {
        alert(response.data);
      })
      .catch(error => {
        alert(error);
      });
  };

  React.useEffect(() => {
    axios
      .get(BASE_URL + 'images')
      .then(response => {
        setImages(response.data);
      })
      .catch(error => {
        console.error(error);
      });
  }, []);
  return (
    <>
      <Helmet>
        <title>Home Page</title>
        <meta name="description" content="A Boilerplate application homepage" />
      </Helmet>
      <div className="flex justify-center">
        <input type="file" multiple onChange={handleCahngeInput} />
        <input
          type="submit"
          className="bg-lime-500 rounded-md text-white px-10 cursor-pointer"
          onClick={handleUploadFile}
        />
      </div>
      <div className="flex justify-center mt-6">
        {images.map(image => {
          return (
            <div className="w-60 h-60">
              <img src={BASE_URL + 'static/' + image.name} alt={image.name} />
            </div>
          );
        })}
      </div>
    </>
  );
}
