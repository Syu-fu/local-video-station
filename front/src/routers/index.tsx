import { type FC } from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import SearchVideoByTagsPage from '../pages/SearchVideoByTagsPage';
import TagCreatePage from '../pages/TagCreatePage';
import VideoCreatePage from '../pages/VideoCreatePage';
import VideoEditPage from '../pages/VideoEditPage';
import VideoPlayerPage from '../pages/VideoPlayerPage';
import VideoListPage from '../pages/list';

const AppRouter: FC = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/list" element={<VideoListPage />} />
        <Route path="/search/video" element={<VideoListPage />} />
        <Route path="video">
          <Route path="tagged" element={<SearchVideoByTagsPage />} />
          <Route path="create" element={<VideoCreatePage />} />
          <Route path="edit">
            <Route path=":id" element={<VideoEditPage />} />
          </Route>
          <Route path=":id" element={<VideoPlayerPage />} />
        </Route>
        <Route path="/tag/create" element={<TagCreatePage />} />
      </Routes>
    </BrowserRouter>
  );
};

export default AppRouter;
