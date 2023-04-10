# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make weighted_least_squares_linear_regressors.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

weighted_least_squares_linear_regressors.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > weighted_least_squares_linear_regressors.tex

weighted_least_squares_linear_regressors.pdf: ../../*.tex ../../trademark.pdf *.tex weighted_least_squares_linear_regressors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode weighted_least_squares_linear_regressors.tex
	../../latexrun weighted_least_squares_linear_regressors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: weighted_least_squares_linear_regressors.pdf
	open weighted_least_squares_linear_regressors.pdf

o:
	make open

terms: weighted_least_squares_linear_regressors.tex
	bbk_sheet -mode ts weighted_least_squares_linear_regressors.tex

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
	rm -f weighted_least_squares_linear_regressors.*
	make reset
	bbk_sheet -mode mk > Makefile