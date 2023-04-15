import { type FC } from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import TagCreatePage from '../pages/TagCreatePage';
import VideoPlayerPage from '../pages/VideoPlayerPage';
import VideoListPage from '../pages/list';

const AppRouter: FC = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/list" element={<VideoListPage />} />
        <Route path="video">
          <Route path=":id" element={<VideoPlayerPage />} />
        </Route>
        <Route path="/tag/create" element={<TagCreatePage />} />
      </Routes>
    </BrowserRouter>
  );
};

export default AppRouter;
