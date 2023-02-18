# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make estimators.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

estimators.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > estimators.tex

estimators.pdf: ../../*.tex ../../trademark.pdf *.tex estimators.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode estimators.tex
	../../latexrun estimators.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: estimators.pdf
	open estimators.pdf

o:
	make open

terms: estimators.tex
	bbk_sheet -mode ts estimators.tex

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
	rm -f estimators.*
	make reset
	bbk_sheet -mode mk > Makefile