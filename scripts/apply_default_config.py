import logging
import socket
import argparse
import os
import toml

def apply_params(home):
    path = "{}/config/config.toml".format(home)
    data = toml.load(path)
    data['p2p']['allow_duplicate_ip'] = True
    data['rpc']['timeout_broadcast_tx_commit'] = '30s'

    f = open(path, 'w')
    toml.dump(data, f)
    f.close()

def main():
    # Setup the logging instance
    logging.basicConfig(level=logging.INFO)

    # Setup the arguments
    parser = argparse.ArgumentParser(description='Apply default parameters to config.toml file')
    parser.add_argument('--home', type=str, help='The path to the chain root folder', required=True)
    args = parser.parse_args()

    apply_params(args.home)
    logging.info("Applied params to config.toml file")

if __name__ == "__main__":
    main()