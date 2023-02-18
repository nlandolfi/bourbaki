# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make generated_topologies.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

generated_topologies.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > generated_topologies.tex

generated_topologies.pdf: ../../*.tex ../../trademark.pdf *.tex generated_topologies.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode generated_topologies.tex
	../../latexrun generated_topologies.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: generated_topologies.pdf
	open generated_topologies.pdf

o:
	make open

terms: generated_topologies.tex
	bbk_sheet -mode ts generated_topologies.tex

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
	rm -f generated_topologies.*
	make reset
	bbk_sheet -mode mk > Makefile