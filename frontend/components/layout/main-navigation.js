import Link from "next/link";
import Logo from "./logo";
import ConnectWallet from "../connect_wallet/connect";
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
        <div className="flex ">
          <div className="px-2 mx-2">
            <ConnectWallet />
          </div>
          <div className="px-2 mx-2">
            <Link href="/peers">
              <a>Peers</a>
            </Link>
          </div>
        </div>
      </nav>
    </header>
  );
}

export default MainNavigation;
