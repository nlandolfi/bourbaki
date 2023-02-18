# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make predictor_performance_metrics.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

predictor_performance_metrics.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > predictor_performance_metrics.tex

predictor_performance_metrics.pdf: ../../*.tex ../../trademark.pdf *.tex predictor_performance_metrics.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode predictor_performance_metrics.tex
	../../latexrun predictor_performance_metrics.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: predictor_performance_metrics.pdf
	open predictor_performance_metrics.pdf

o:
	make open

terms: predictor_performance_metrics.tex
	bbk_sheet -mode ts predictor_performance_metrics.tex

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
	rm -f predictor_performance_metrics.*
	make reset
	bbk_sheet -mode mk > Makefile