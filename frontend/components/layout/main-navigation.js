import Link from "next/link";
import Logo from "./logo";
import classes from "./main-navigation.module.css";

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
          <Link href="/peers">
            <a>Peers</a>
          </Link>
        </ul>
      </nav>
    </header>
  );
}

export default MainNavigation;
