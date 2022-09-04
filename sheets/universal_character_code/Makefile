# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make universal_character_code.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

universal_character_code.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > universal_character_code.tex

universal_character_code.pdf: ../../*.tex ../../trademark.pdf *.tex universal_character_code.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode universal_character_code.tex
	../../latexrun universal_character_code.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: universal_character_code.pdf
	open universal_character_code.pdf

o:
	make open

terms: universal_character_code.tex
	bbk_sheet -mode ts universal_character_code.tex

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
	rm -f universal_character_code.*
	make reset
	bbk_sheet -mode mk > Makefile