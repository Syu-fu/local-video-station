import type Tag from './tag';

type VideoFormInputs = {
  id: string;
  title: string;
  titleReading: string;
  thumbnail: File;
  video: File;
  tags: Tag[];
};

export default VideoFormInputs;
