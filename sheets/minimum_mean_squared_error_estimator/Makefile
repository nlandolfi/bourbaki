# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make minimum_mean_squared_error_estimator.pdf

minimum_mean_squared_error_estimator.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > minimum_mean_squared_error_estimator.tex

minimum_mean_squared_error_estimator.pdf: ../../*.tex ../../trademark.pdf *.tex minimum_mean_squared_error_estimator.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode minimum_mean_squared_error_estimator.tex
	../../latexrun minimum_mean_squared_error_estimator.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: minimum_mean_squared_error_estimator.pdf
	open minimum_mean_squared_error_estimator.pdf

o:
	make open

terms: minimum_mean_squared_error_estimator.tex
	bbk_sheet -mode ts minimum_mean_squared_error_estimator.tex

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
	rm -f minimum_mean_squared_error_estimator.*
	make reset
	bbk_sheet -mode mk > Makefile