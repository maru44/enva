import { NextPage } from 'next'
import Link from 'next/link'
import { loginUrl } from '../../../config/aws'
import { PageProps } from '../../../types/page'

const LoginPage2: NextPage<PageProps> = (props) => {
  return (
    <div>
      <div>
        <h2>Sign in</h2>
        <Link href={loginUrl} passHref>
          <a>Sign in</a>
        </Link>
      </div>
    </div>
  )
}

export default LoginPage2
