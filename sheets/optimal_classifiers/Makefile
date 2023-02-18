# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make optimal_classifiers.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

optimal_classifiers.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > optimal_classifiers.tex

optimal_classifiers.pdf: ../../*.tex ../../trademark.pdf *.tex optimal_classifiers.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode optimal_classifiers.tex
	../../latexrun optimal_classifiers.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: optimal_classifiers.pdf
	open optimal_classifiers.pdf

o:
	make open

terms: optimal_classifiers.tex
	bbk_sheet -mode ts optimal_classifiers.tex

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
	rm -f optimal_classifiers.*
	make reset
	bbk_sheet -mode mk > Makefile