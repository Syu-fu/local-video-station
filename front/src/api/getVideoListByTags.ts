import axios from '../lib/axios';
import type Tag from '../types/tag';
import type Video from '../types/video';

const getVideoListByTags = async (tag: Tag[], page = 1): Promise<Video[]> => {
  const tagIds = tag.map((tag) => tag.id).join(',');
  const response = await axios.get<Video[]>(
    `video/list?page=${page}&tags=${tagIds}`
  );

  return response.data;
};

export default getVideoListByTags;
