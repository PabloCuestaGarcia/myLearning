

from st2common.runners.base_action import BaseAction


class BaseAction(Action):
    def run(self, **kwargs):
        pass

    def __init__(self, config):

    