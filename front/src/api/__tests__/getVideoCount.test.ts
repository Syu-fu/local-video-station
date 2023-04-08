import { expect } from 'chai';
import { rest } from 'msw';
import { setupServer } from 'msw/node';
import getVideoCount from '../getVideoCount';

const server = setupServer(
  rest.get('http://localhost:5173/video/count', async (_, res, ctx) => {
    return await res(ctx.json({ count: 10 }));
  })
);

beforeAll(() => {
  server.listen();
});
afterEach(() => {
  server.resetHandlers();
});
afterAll(() => {
  server.close();
});

describe('getVideoCount', () => {
  it('fetches the count of videos', async () => {
    const promise = getVideoCount();
    const count = await promise;
    expect(count).to.deep.equal({ count: 10 });
  });
});
