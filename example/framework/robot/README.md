## Usage

### Robot环境配置

基于Debian系统默认已经安装Python3了

```
apt install -y python3-apt python3-pip
python3 -m venv /tmp/venv
source /tmp/venv/bin/activate
pip install robotframework
# 验证
robot --version
```

### Robot测试用例编写

```
ztf robot -p 1 --verbose robot -d results testcase
```

登录禅道，测试-用例-单元测试(默认是所有类型)

