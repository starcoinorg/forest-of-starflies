import React, { useState, useCallback, useEffect } from "react";
import classnames from "classnames";
import { providers } from "@starcoin/starcoin";
import StarMaskOnboarding from "@starcoin/starmask-onboarding";
import Router from "next/router";
import Link from "next/link";

const { isStarMaskInstalled } = StarMaskOnboarding;

export default function ConnectWallet() {
  const [isStarMaskConnected, setConnected] = useState(false);
  // 已连接账户
  const [account, setAccount] = useState([]);
  const [isInstall, setInstall] = useState(true);
  const freshConnected = useCallback(async () => {
    const newAccounts = await window.starcoin.request({
      method: "stc_requestAccounts",
    });
    setAccount([...newAccounts]);
    setConnected(newAccounts && newAccounts.length > 0);
  }, []);

  useEffect(() => {
    if (!isStarMaskInstalled()) {
      setInstall(false);
      alert("没有安装 starmask 插件！");
      onboarding.startOnboarding();
    } else {
      setInstall(true);
    }
  }, [freshConnected]);

  useEffect(() => {
    try {
      starcoinProvider = new providers.Web3Provider(window.starcoin, "any");
    } catch {
      setInstall(false);
    }
  }, []);

  const handleClick = useCallback(() => {
    if (isStarMaskConnected) {
      if (onboarding) {
        onboarding.stopOnboarding();
      }
    } else {
      freshConnected();
    }
  }, [freshConnected, isStarMaskConnected]);

  useEffect(() => {
    if (isStarMaskConnected) {
      Router.push("/claimlist?address=" + account[0]);
    }
  }, [isStarMaskConnected, account]);

  return (
    <div
      className={isStarMaskConnected ? "cursor-not-allowed" : "cursor-pointer"}
      onClick={!isStarMaskConnected ? handleClick : null}
    >
      <div>{isStarMaskConnected ? "Connected" : "Connect"}</div>
      <div>
        {isStarMaskConnected ? (
          <Link
            href={{ pathname: "/claimlist", query: { address: account[0] } }}
          >
            <a className="cursor-pointer">{account}</a>
          </Link>
        ) : null}
      </div>
    </div>
  );
}
