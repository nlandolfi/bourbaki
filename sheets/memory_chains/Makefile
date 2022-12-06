# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make memory_chains.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

memory_chains.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > memory_chains.tex

memory_chains.pdf: ../../*.tex ../../trademark.pdf *.tex memory_chains.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode memory_chains.tex
	../../latexrun memory_chains.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: memory_chains.pdf
	open memory_chains.pdf

o:
	make open

terms: memory_chains.tex
	bbk_sheet -mode ts memory_chains.tex

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
	rm -f memory_chains.*
	make reset
	bbk_sheet -mode mk > Makefile