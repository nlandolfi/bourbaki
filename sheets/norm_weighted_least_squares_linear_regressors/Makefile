# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make norm_weighted_least_squares_linear_regressors.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

norm_weighted_least_squares_linear_regressors.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > norm_weighted_least_squares_linear_regressors.tex

norm_weighted_least_squares_linear_regressors.pdf: ../../*.tex ../../trademark.pdf *.tex norm_weighted_least_squares_linear_regressors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode norm_weighted_least_squares_linear_regressors.tex
	../../latexrun norm_weighted_least_squares_linear_regressors.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: norm_weighted_least_squares_linear_regressors.pdf
	open norm_weighted_least_squares_linear_regressors.pdf

o:
	make open

terms: norm_weighted_least_squares_linear_regressors.tex
	bbk sheet -mode ts norm_weighted_least_squares_linear_regressors.tex

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
	rm -f norm_weighted_least_squares_linear_regressors.*
	make reset
	bbk sheet -mode mk > Makefile