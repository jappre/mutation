#!/usr/bin/env python
# -*- coding: utf-8 -*-
import config_default
from config_override import configs

configs = config_default.configs

try:
    import config_override
    configs = merge(configs, config_override.configs)
except ImportError:
    pass
