# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make topological_vector_spaces.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

topological_vector_spaces.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > topological_vector_spaces.tex

topological_vector_spaces.pdf: ../../*.tex ../../trademark.pdf *.tex topological_vector_spaces.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode topological_vector_spaces.tex
	../../latexrun topological_vector_spaces.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: topological_vector_spaces.pdf
	open topological_vector_spaces.pdf

o:
	make open

terms: topological_vector_spaces.tex
	bbk_sheet -mode ts topological_vector_spaces.tex

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
	rm -f topological_vector_spaces.*
	make reset
	bbk_sheet -mode mk > Makefile