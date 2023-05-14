import axios from '../lib/axios';
import type Video from '../types/video';
import type VideoFormInputs from '../types/videoFormInputs';

const putVideoDetail = async (video: VideoFormInputs): Promise<Video> => {
  console.log(video);
  const formData = new FormData();
  formData.append('id', video.id);
  formData.append('title', video.title);
  formData.append('titleReading', video.titleReading);
  formData.append('thumbnail', video.thumbnail);
  formData.append('video', video.video);
  formData.append('tags', JSON.stringify(video.tags));

  const response = await axios.put<Video>('video', formData);

  return response.data;
};

export default putVideoDetail;
