# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make decision_problems.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

decision_problems.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > decision_problems.tex

decision_problems.pdf: ../../*.tex ../../trademark.pdf *.tex decision_problems.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode decision_problems.tex
	../../latexrun decision_problems.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: decision_problems.pdf
	open decision_problems.pdf

o:
	make open

terms: decision_problems.tex
	bbk_sheet -mode ts decision_problems.tex

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
	rm -f decision_problems.*
	make reset
	bbk_sheet -mode mk > Makefile