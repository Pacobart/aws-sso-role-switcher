setup:
	pip install virtualenv
	virtualenv venv
	pip install -r requirements.txt

clean:
	virtualenv --clear venv

dev:
	. ./venv/bin/activate

install:
	python3 -m pip install -e .

uninstall:
	python3 -m pip uninstall -e .

run:
	#. aws-sso-role-switcher.sh -v
	#. aws-sso-role-switcher.sh -h
	#. aws-sso-role-switcher.sh
	python bin/aws-sso-role-switcher

test:
	python tests/test.py

package:
	python3 -m pip install --user --upgrade twine
	python3 setup.py sdist bdist_wheel
	python3 -m twine upload dist/*