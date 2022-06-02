import Link from "next/link";
import Logo from "./logo";
import classes from './main-navigation.module.css'

function MainNavigation() {
  return (
    <header className={classes.header}>
      <Link href="/">
        <a>
          <Logo></Logo>
        </a>
      </Link>
      <nav>
          <ul>
              <a>
                  <Link href='/peers'>Peers</Link>
              </a>
          </ul>
      </nav>
    </header>
  );
}

export default MainNavigation;