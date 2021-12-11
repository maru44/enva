import { NextPage } from 'next'
import { fetchCreatePost } from '../../../http/post'
import { PageProps } from '../../../types/page'
import { PostInput } from '../../../types/post'

const PostCreate: NextPage<PageProps> = (props) => {
  const startPost = async () => {
    const input: PostInput = {
      title: 'test',
      abstract: 'this is a test',
      content: 'koreha test desu\n\nchuui shitene',
    }

    const res = await fetchCreatePost(input)
    console.log(res)
    const ret = await res.json()

    console.log(ret)
  }

  return (
    <div>
      <div>
        <button type="button" onClick={startPost}>
          Post Test
        </button>
      </div>
    </div>
  )
}

export default PostCreate
