# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make unbiased_estimators.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

unbiased_estimators.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > unbiased_estimators.tex

unbiased_estimators.pdf: ../../*.tex ../../trademark.pdf *.tex unbiased_estimators.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode unbiased_estimators.tex
	../../latexrun unbiased_estimators.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: unbiased_estimators.pdf
	open unbiased_estimators.pdf

o:
	make open

terms: unbiased_estimators.tex
	bbk_sheet -mode ts unbiased_estimators.tex

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
	rm -f unbiased_estimators.*
	make reset
	bbk_sheet -mode mk > Makefile