import type Tag from './tag';

type VideoFormInputs = {
  title: string;
  titleReading: string;
  thumbnail: File;
  video: File;
  tags: Tag[];
};

export default VideoFormInputs;
