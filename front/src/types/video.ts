import type Tag from './tag';

type Video = {
  id: string;
  title: string;
  titleReading: string;
  url: string;
  thumbnailUrl: string;
  tags: Tag[];
};

export default Video;
