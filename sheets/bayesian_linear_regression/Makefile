# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make bayesian_linear_regression.pdf

bayesian_linear_regression.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > bayesian_linear_regression.tex

bayesian_linear_regression.pdf: ../../*.tex ../../trademark.pdf *.tex bayesian_linear_regression.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode bayesian_linear_regression.tex
	../../latexrun bayesian_linear_regression.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: bayesian_linear_regression.pdf
	open bayesian_linear_regression.pdf

o:
	make open

terms: bayesian_linear_regression.tex
	bbk_sheet -mode ts bayesian_linear_regression.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f bayesian_linear_regression.*
	make reset
	bbk_sheet -mode mk > Makefile